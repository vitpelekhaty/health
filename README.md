# health / checkers

Custom health checkers for [github.com/InVisionApp/go-health](github.com/InVisionApp/go-health)

## Usage

```go
import check "https://github.com/vitpelekhaty/health/checkers"
```

## MongoDB health check

A standard MongoDB's health checker uses a driver [mgo](github.com/globalsign/mgo). This a health checker uses [an official MongoDB driver](https://github.com/mongodb/mongo-go-driver) for Go.

### Example

```go
import (
    "github.com/InVisionApp/go-health"
    check "https://github.com/vitpelekhaty/health/checkers"
)

func HealthCheck(mongo *mongo.Client) (*health.Health, error) {
    healthCheck := health.New()

    mongoCheck, err := check.NewMongo(&MongoConfig{
        Pinger: mongo,
    })

    if err != nil {
        return nil, err
    }

    err = healthCheck.AddChecks(
        []*health.Config{
            &health.Config{
                Name:     "mongo-check",
                Checker:  mongoCheck,
                Interval: time.Duration(3) * time.Second,
            },
        },
    )

    if err != nil {
        return nil, err
    }

    return healthCheck, nil
}
```
