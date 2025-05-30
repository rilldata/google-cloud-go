// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pubsub_test

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/iterator"
)

func ExampleClient_Subscriptions() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		// TODO: Handle error.
	}
	// List all subscriptions of the project.
	it := client.Subscriptions(ctx)
	_ = it // TODO: iterate using Next.
}

func ExampleSubscriptionIterator_Next() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		// TODO: Handle error.
	}
	// List all subscriptions of the project.
	it := client.Subscriptions(ctx)
	for {
		sub, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		fmt.Println(sub)
	}
}
