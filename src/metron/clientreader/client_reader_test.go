package clientreader_test

import (
	"doppler/dopplerservice"
	"metron/clientreader"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	preferredProtocol string
	mockPool          *mockClientPool
	event             dopplerservice.Event
	clientPool        map[string]clientreader.ClientPool
)

var _ = Describe("clientreader", func() {

	JustBeforeEach(func() {
		clientPool[preferredProtocol] = mockPool
	})

	BeforeEach(func() {
		mockPool = newMockClientPool()
		clientPool = make(map[string]clientreader.ClientPool)
	})

	Describe("Read", func() {
		Context("TLS PreferredProtocol", func() {
			BeforeEach(func() {
				preferredProtocol = "tls"
			})

			It("doesn't panic if there are tls dopplers", func() {
				event = dopplerservice.Event{
					UDPDopplers: []string{},
					TLSDopplers: []string{"10.0.0.1"}}

				l := len(event.TLSDopplers)
				mockPool.SetAddressesOutput.ret0 <- l
				clientPool := map[string]clientreader.ClientPool{"tls": mockPool}
				Expect(func() {
					clientreader.Read(clientPool, []string{preferredProtocol}, event)
				}).ToNot(Panic())
				Eventually(mockPool.SetAddressesCalled).Should(Receive())
			})

			It("panics if there are no tls dopplers", func() {
				event = dopplerservice.Event{
					UDPDopplers: []string{"1.1.1.1.", "2.2.2.2"},
					TLSDopplers: []string{},
				}
				l := len(event.TLSDopplers)
				mockPool.SetAddressesOutput.ret0 <- l

				Expect(func() { clientreader.Read(clientPool, []string{preferredProtocol}, event) }).To(Panic())
				Eventually(mockPool.SetAddressesCalled).Should(Receive())
			})

		})
		Context("UDP PreferredProtocol", func() {
			BeforeEach(func() {
				preferredProtocol = "udp"
			})
			It("doesn't panic for udp only dopplers", func() {
				event := dopplerservice.Event{
					UDPDopplers: []string{"10.0.0.1"},
					TLSDopplers: []string{},
				}
				l := len(event.UDPDopplers)
				mockPool.SetAddressesOutput.ret0 <- l

				Expect(func() { clientreader.Read(clientPool, []string{preferredProtocol}, event) }).ToNot(Panic())
				Eventually(mockPool.SetAddressesCalled).Should(Receive())
			})

		})

		Context("TCP PreferredProtocol", func() {
			BeforeEach(func() {
				preferredProtocol = "tcp"
			})

			It("doesn't panic if there are tcp dopplers", func() {
				event = dopplerservice.Event{
					UDPDopplers: []string{},
					TLSDopplers: []string{"10.0.0.1"},
					TCPDopplers: []string{"10.0.0.1"},
				}

				l := len(event.TCPDopplers)
				mockPool.SetAddressesOutput.ret0 <- l
				Expect(func() { clientreader.Read(clientPool, []string{preferredProtocol}, event) }).ToNot(Panic())
				Eventually(mockPool.SetAddressesCalled).Should(Receive())
			})

			It("panics if there are no tcp dopplers", func() {
				event = dopplerservice.Event{
					UDPDopplers: []string{},
					TLSDopplers: []string{"10.0.0.1"},
				}

				l := len(event.TCPDopplers)
				mockPool.SetAddressesOutput.ret0 <- l
				Expect(func() { clientreader.Read(clientPool, []string{preferredProtocol}, event) }).To(Panic())
				Eventually(mockPool.SetAddressesCalled).Should(Receive())

			})
		})
	})
})
