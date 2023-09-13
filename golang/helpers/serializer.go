package helpers

func SerializeData(data interface{}) map[string]interface{} {
    return map[string]interface{}{
        "data": data,
    }
}

func SerializeError(message interface{}, status int) map[string]interface{} {
    return map[string]interface{}{
        "error": map[string]interface{}{
            "message": message,
            "status": status,
        },
    }
}
