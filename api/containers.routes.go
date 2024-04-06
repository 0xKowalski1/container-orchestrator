package api

import (
	"net/http"

	"0xKowalski1/container-orchestrator/models"
	statemanager "0xKowalski1/container-orchestrator/state-manager"

	"github.com/gin-gonic/gin"
)

// GET /containers
func getContainers(c *gin.Context, _statemanager *statemanager.StateManager) {
	containers, err := _statemanager.ListContainers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"containers": containers,
	})

}

// POST /containers
func createContainer(c *gin.Context, _statemanager *statemanager.StateManager) {
	var req models.CreateContainerRequest
	// Parse the JSON body to the CreateContainerRequest struct.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := _statemanager.AddContainer(models.Container{ID: req.ID, Image: req.Image, Env: req.Env, StopTimeout: req.StopTimeout})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createdContainer := models.Container{
		ID:          req.ID,
		Image:       req.Image,
		Env:         req.Env,
		StopTimeout: req.StopTimeout,
	}

	c.JSON(http.StatusCreated, gin.H{
		"container": createdContainer,
	})
}

// PATCH /containers
func updateContainer(c *gin.Context, _statemanager *statemanager.StateManager) {
	containerID := c.Param("id")

	var req models.UpdateContainerRequest
	// Parse the JSON body to the UpdateContainerRequest struct.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := _statemanager.PatchContainer(containerID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

// GET /containers/{id}
func getContainer(c *gin.Context, _statemanager *statemanager.StateManager) {
	containerID := c.Param("id")

	container, err := _statemanager.GetContainer(containerID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Container not found"})
		return
	}

	c.JSON(http.StatusOK, container)
}

// DELETE /containers/{id}
func deleteContainer(c *gin.Context, _statemanager *statemanager.StateManager) {
	containerID := c.Param("id")

	// Should mark for deletion!
	err := _statemanager.RemoveContainer(containerID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "true"})
}

// POST /containers/{id}/start
func startContainer(c *gin.Context, _statemanager *statemanager.StateManager) {
	containerID := c.Param("id")

	desiredStatus := "running"

	err := _statemanager.PatchContainer(containerID, models.UpdateContainerRequest{
		DesiredStatus: &desiredStatus,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Container starting"})
}

// POST /containers/{id}/stop
func stopContainer(c *gin.Context, _statemanager *statemanager.StateManager) {
	containerID := c.Param("id")

	desiredStatus := "stopped"

	err := _statemanager.PatchContainer(containerID, models.UpdateContainerRequest{
		DesiredStatus: &desiredStatus,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Container stopping"})
}

// GET /containers/{id}/logs
func getContainerLogs(c *gin.Context, _statemanager *statemanager.StateManager) {
	// containerID := c.Param("id") // Retrieve the container ID from the URL parameter.
}
