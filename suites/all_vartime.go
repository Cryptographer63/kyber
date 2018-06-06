// +build vartime

package suites

import (
	"github.com/dedis/kyber/group/curve25519"
	"github.com/dedis/kyber/group/nist"
	"github.com/dedis/kyber/pairing/bn256"
)

func init() {
	register(curve25519.NewBlakeSHA256Curve25519(false))
	register(curve25519.NewBlakeSHA256Curve25519(true))
	register(nist.NewBlakeSHA256P256())
	register(nist.NewBlakeSHA256QR512())
	register(bn256.NewSuite().G1().(Suite))
	register(bn256.NewSuite().G2().(Suite))
	register(bn256.NewSuite().GT().(Suite))
}
