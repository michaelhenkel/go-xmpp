package main

import (
	"github.com/bdlm/log"

	"github.com/michaelhenkel/go-xmpp"
	"github.com/michaelhenkel/go-xmpp/stanza"
)

func joinMUC(c xmpp.Sender, toJID *stanza.Jid) error {
	return c.Send(stanza.Presence{Attrs: stanza.Attrs{To: toJID.Full()},
		Extensions: []stanza.PresExtension{
			stanza.MucPresence{
				History: stanza.History{MaxStanzas: stanza.NewNullableInt(0)},
			}},
	})
}

func leaveMUCs(c xmpp.Sender, mucsToLeave []*stanza.Jid) {
	for _, muc := range mucsToLeave {
		if err := c.Send(stanza.Presence{Attrs: stanza.Attrs{
			To:   muc.Full(),
			Type: stanza.PresenceTypeUnavailable,
		}}); err != nil {
			log.WithField("muc", muc).Errorf("error on leaving muc: %s", err)
		}
	}
}
