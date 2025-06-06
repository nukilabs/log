package main

import (
	"os"

	"github.com/nukilabs/log"
)

func main() {
	logger := log.New(os.Stdout, log.WithIndex(1), log.WithPrefix("test"), log.WithLevel(log.DebugLevel))
	logger.Hint("Loading login page.")
	logger.Debug("Solving funcaptcha.", "step", 1)
	logger.Debug("Solving funcaptcha.", "step", 2)
	logger.Error("Failed to solve funcaptcha.")
	logger.Debug("Solving funcaptcha.", "step", 1)
	logger.Debug("Solving funcaptcha.", "step", 2)
	logger.Debug("Solving funcaptcha.", "step", 3)
	logger.Info("Login successful.")
	logger.Hint("Loading event page.")
	logger.Warn("Out of stock.")
	logger.Warn("Out of stock.")
	logger.Warn("Out of stock.")
	logger.Miss("Seats are already taken.")
	logger.Warn("Out of stock.")
	logger.Warn("Out of stock.")
	logger.Cart("Added to cart.")
	logger.Hint("Loading checkout page.")
	logger.Debug("Submitting address.")
	logger.Debug("Submitting payment.")
	logger.Debug("Submitting shipping.")
	logger.Done("Checkout successful.")
}
