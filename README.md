
# logit

An easy level log printer with line number.

## Example
```
package main

import (
	"github.com/winary/logit"
)

var log = logit.NewLogPkg("your-pkg")

func main() {
	// logit.UseDefaultWriter(logit.LevelInfo, "worker.log")

	log.Debug("%d == %s", 1, "1")
	log.Info("%d == %s", 2, "2")

}
```
### Output
```
time=2020-06-02 11:23:22.969, level=debug, file=your-pkg/logit.go:12, 1 == 1
time=2020-06-02 11:23:22.979, level=info , file=your-pkg/logit.go:13, 2 == 2
```
