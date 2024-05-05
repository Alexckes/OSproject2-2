package builtins

import (
"fmt"
"time"
)


func Time() error {
	fmt.Printf( "The current time is:" + time.Now().String());
	return fmt.Errorf("");
}