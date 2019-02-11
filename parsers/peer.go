package parsers

import (
	"fmt"
	"strconv"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type Peer struct {
	data []types.Peer
}

func (l *Peer) parse(line string, parts []string, comment string) (*types.Peer, error) {
	if len(parts) >= 2 {
		adr := common.StringSplitIgnoreEmpty(parts[2], ':')
		if len(adr) >= 2 {
			if port, err := strconv.ParseInt(adr[1], 10, 64); err == nil {
				return &types.Peer{
					Name:    parts[1],
					IP:      adr[0],
					Port:    port,
					Comment: comment,
				}, nil
			}
		}
	}
	return nil, &errors.ParseError{Parser: "PeerLines", Line: line}
}

func (l *Peer) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(l.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(l.data))
	for index, peer := range l.data {
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("peer %s %s:%d", peer.Name, peer.IP, peer.Port),
			Comment: peer.Comment,
		}
	}
	return result, nil
}