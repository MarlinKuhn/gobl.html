package pdf

import (
	"context"

	"github.com/invopop/princepdf"
)

type princeConvertor struct {
	client *princepdf.Client
}

func newPrinceConvertor() (*princeConvertor, error) {
	c := &princeConvertor{client: princepdf.New()}

	if err := c.client.Start(); err != nil {
		return nil, err
	}

	return c, nil
}

func (pc *princeConvertor) HTML(ctx context.Context, data []byte, opts ...Option) ([]byte, error) {
	o := prepareOptions(opts)
	j := new(princepdf.Job)
	j.Input = &princepdf.Input{
		Src: "data.html",
	}
	j.Files = map[string][]byte{
		"data.html": data,
	}

	if len(o.attachments) > 0 {
		for _, a := range o.attachments {
			j.Files[a.Filename] = a.Data
			if j.PDF == nil {
				j.PDF = new(princepdf.PDF)
			}
			j.PDF.Attach = append(j.PDF.Attach, &princepdf.Attachment{
				URL:         a.Filename,
				Filename:    a.Filename,
				Description: a.Description,
			})
		}
	}

	if o.metadata != nil {
		j.Metadata = &princepdf.Metadata{
			Title:    o.metadata.Title,
			Subject:  o.metadata.Subject,
			Author:   o.metadata.Author,
			Keywords: o.metadata.Keywords,
			Creator:  o.metadata.Creator,
		}
	}

	return pc.client.Run(j)
}