package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	xyz "github.com/restlabs/pulumi-routeros/sdk/go/routeros"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := xyz.NewRandom(ctx, "myRandomResource", &xyz.RandomArgs{
			Length: pulumi.Int(24),
		})
		if err != nil {
			return err
		}

		return nil
	})
}
