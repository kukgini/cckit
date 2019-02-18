package encryption

import (
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
	"github.com/s7techlab/cckit/response"
	"github.com/s7techlab/cckit/router"
)

// ArgsDecryptIfKeyProvided  - pre middleware, decrypts chaincode method arguments if key provided in transient map
func ArgsDecryptIfKeyProvided(next router.ContextHandlerFunc, pos ...int) router.ContextHandlerFunc {
	return argsDecryptor(next, false)
}

// ArgsDecryptIfKeyProvided  - pre middleware, decrypts chaincode method arguments, key must be provided in transient map
func ArgsDecrypt(next router.ContextHandlerFunc, pos ...int) router.ContextHandlerFunc {
	return argsDecryptor(next, true)
}

func decryptReplaceArgs(key []byte, c router.Context) error {
	args, err := DecryptArgs(key, c.Stub().GetArgs())
	if err != nil {
		return errors.Wrap(err, `args`)
	}
	c.ReplaceArgs(args)
	return nil
}

func argsDecryptor(next router.ContextHandlerFunc, keyShouldBeProvided bool) router.ContextHandlerFunc {
	return func(c router.Context) peer.Response {
		key, err := KeyFromTransient(c)
		// no key provided

		if err != nil {
			c.Logger().Debugf(`no decrypt key provided: %s`, err)
			if err == ErrKeyNotDefinedInTransientMap && keyShouldBeProvided {
				return response.Error(err)
			}
			return next(c)
		}

		if err = decryptReplaceArgs(key, c); err != nil {
			return response.Error(err)
		}

		return next(c)
	}
}

// EncStateContext replaces default state with encrypted state
func EncStateContext(next router.HandlerFunc, pos ...int) router.HandlerFunc {
	return func(c router.Context) (interface{}, error) {
		s, err := StateWithTransientKey(c)
		if err != nil {
			return nil, err
		}
		c.UseState(s)
		//cc := &EncryptedStateContext{c}
		return next(c)
	}
}
