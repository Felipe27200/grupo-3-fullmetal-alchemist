package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"alchemy-system/database"
	"alchemy-system/models"
	"alchemy-system/queue"
	"alchemy-system/services"
)

func StartTransmutationWorker() {
	for {
		// Esperar trabajo en Redis (bloqueante)
		data, err := queue.RedisClient.BRPop(context.Background(), 0*time.Second, "transmutation_queue").Result()
		if err != nil {
			fmt.Println("Worker error reading queue:", err)
			continue
		}

		// Data[1] contiene el JSON enviado
		var job map[string]any
		if err := json.Unmarshal([]byte(data[1]), &job); err != nil {
			fmt.Println("Worker JSON error:", err)
			continue
		}

		fmt.Println("Worker received job:", job)

		var transmutation models.Transmutation
		if err := database.DB.First(&transmutation, job["id"]).Error; err != nil {
			fmt.Println("Transmutation not found:", err)
			continue
		}

		// Marcar como en proceso
		transmutation.Status = "processing"
		database.DB.Save(&transmutation)

		// Simular proceso real
		time.Sleep(2 * time.Second)

		// Aquí puedes implementar lógica real de transmutación
		transmutation.Approved = true
		transmutation.Result = fmt.Sprintf("Processed: %s", transmutation.Input)
		transmutation.Status = "completed"

		now := time.Now()
		transmutation.ExecutedAt = &now

		if err := database.DB.Save(&transmutation).Error; err != nil {
			fmt.Println("Worker DB save error:", err)
			continue
		}

		// Registrar auditoría
		services.CreateAudit(models.Audit{
			Action:   "COMPLETE",
			Entity:   "transmutation",
			EntityID: transmutation.ID,
			Message:  "Transmutation processed successfully by worker",
		})

		fmt.Println("Transmutation completed:", transmutation.ID)
	}
}
