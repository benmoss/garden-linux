package performance_test

import (
	"runtime"
	"strconv"

	"log"
	"os"

	"github.com/cloudfoundry-incubator/garden"
	gclient "github.com/cloudfoundry-incubator/garden/client"
	"github.com/cloudfoundry-incubator/garden/client/connection"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	creates       = 40
	createSamples = 5
)

var _ = Describe("Concurrent container creation", func() {

	BeforeEach(func() {
		runtime.GOMAXPROCS(runtime.NumCPU())
	})

	Measure("multiple concurrent creates", func(b Benchmarker) {
		if os.Getenv("GARDEN_PERFORMANCE") == "" {
			log.Println("GARDEN_PERFORMANCE undefined; skipping")
			return
		}
		handles := []string{}

		b.Time("concurrent creations", func() {
			chans := []chan string{}
			for i := 0; i < creates; i++ {
				ch := make(chan string, 1)
				go func(c chan string, index int) {
					defer GinkgoRecover()
					client := gclient.New(connection.New("tcp", "localhost:7777"))
					b.Time("create-"+strconv.Itoa(index), func() {
						ctr, err := client.Create(garden.ContainerSpec{})
						Expect(err).ToNot(HaveOccurred())
						c <- ctr.Handle()
					})
				}(ch, i)
				chans = append(chans, ch)
			}

			for _, ch := range chans {
				handle := <-ch
				if handle != "" {
					handles = append(handles, handle)

				}
			}
		})

		for _, handle := range handles {
			client := gclient.New(connection.New("tcp", "localhost:7777"))
			Expect(client.Destroy(handle)).To(Succeed())
		}

	}, createSamples)

})
