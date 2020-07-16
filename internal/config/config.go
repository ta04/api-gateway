/*
Dear Programmers,

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*                                                 *
*	This file belongs to Kevin Veros Hamonangan   *
*	and	Fandi Fladimir Dachi and is a part of     *
*	our	last project as the student of Del        *
*	Institute of Technology, Sitoluama.           *
*	Please contact us via Instagram:              *
*	sleepingnext and fandi_dachi                  *
*	before copying this file.                     *
*	Thank you, buddy. 😊                          *
*                                                 *
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*/

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from .env into the system
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

// MicroWebName will return MICRO_WEB_NAME from .env file
func MicroWebName() string {
	return lookupEnv("MICRO_WEB_NAME")
}

// MicroWebPort will return MICRO_WEB_PORT from .env file
func MicroWebPort() string {
	return lookupEnv("MICRO_WEB_PORT")
}

// SecretKey will return SECRET_KEY from .env file
func SecretKey() string {
	return lookupEnv("SECRET_KEY")
}

func lookupEnv(key string) string {
	env, exist := os.LookupEnv(key)
	if !exist {
		log.Fatal(key, "is not exists in .env file")
	}

	return env
}
