package domain

import "errors"

var ErrGroupNotFound = errors.New("group not found")

var ErrGroupConflict = errors.New("group already exists")
