package fdb

import (
    "os"
)

func (c *Collection) KeyExists(key string) bool {
    _, err := os.Stat(c.Key2IndexPath(key))

    if os.IsNotExist(err) {
        return false
    }

    if err != nil {
        panic(err)
    }

    return true
}
