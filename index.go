package main

import (
	"database/sql"
	"kimb/internal/database"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("[FATAL] No port configured >>>")
	}

	databaseURL := os.Getenv("POSTGRE")
	if databaseURL == "" {
		log.Fatal("[FATAL] No database configured >>>")
	}

	con, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("[FATAL] Error connecting to database >>>\n", err)
	}

	apiCfg := apiConfig{
		DB: database.New(con),
	}

	/**ROUTING */
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	/**ROUTER */
	v1router := chi.NewRouter()
	v1router.Get("/status", handlerReady)
	v1router.Get("/error", handlerError)
	/**RESGISTER / LOGIN*/
	v1router.Post("/register", apiCfg.handlerUserCreate)
	v1router.Post("/login", apiCfg.handlerUserLogin)
	/**USERS */
	v1router.Get("/users/{user}", apiCfg.handlerGetUserFromNameOrID)
	v1router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUserFromAPI))
	v1router.Get("/users_likes", apiCfg.handlerGetAllUserLikes)
	/**FOLLOWS */
	v1router.Post("/users_follows", apiCfg.middlewareAuth(apiCfg.handlerFollow))
	v1router.Delete("/users_follows/{followedID}", apiCfg.middlewareAuth(apiCfg.handlerUnfollow))
	v1router.Get("/users_follows", apiCfg.handlerGetFollowing)
	v1router.Get("/users_followers", apiCfg.handlerGetFollowers)
	/**POSTS */
	v1router.Get("/posts", apiCfg.handlerGetPosts)
	v1router.Post("/posts", apiCfg.middlewareAuth(apiCfg.handlerCreatePost))
	v1router.Get("/posts/{userID}", apiCfg.handlerGetPostsFromUser)
	v1router.Delete("/posts/{postID}", apiCfg.middlewareAuth(apiCfg.handlerDeletePost))
	v1router.Post("/posts_likes", apiCfg.middlewareAuth(apiCfg.handlerLikePost))
	v1router.Get("/posts_likes", apiCfg.handlerGetPostLikes)
	v1router.Delete("/posts_likes/{postID}", apiCfg.middlewareAuth(apiCfg.handlerRemovePostLike))
	/**MOUNT */
	router.Mount("/v1", v1router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("| [INFO] Server started at port %v", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("[FATAL] error stopped with error >>>\n", err)
	}
}
