package middlewares

// func AuthorizationMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Extract user from request context
// 		user, err := GetUserFromContext(r.Context())
// 		if err != nil {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		// Check if user has the required permissions
// 		if !user.HasPermission("required_permission") {
// 			http.Error(w, "Forbidden", http.StatusForbidden)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }
