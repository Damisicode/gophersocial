package main

import (
	"net/http"
)

// HealthCheck godoc
//
//	@Summary		health check
//	@Description	check health status of the server
//	@Tags			Ops
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int		true	"Post ID"
//	@Success		200	{object}	string	"server is ok"
//	@Failure		404	{object}	error	"page not found"
//	@Failure		500	{object}	error	"Server not available"
//	@Security		ApiKeyAuth
//	@Router			/health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}

	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
	}
}
