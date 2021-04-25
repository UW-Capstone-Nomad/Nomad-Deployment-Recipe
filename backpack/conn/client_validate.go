package conn

import (
	"github.com/hashicorp/nomad/api"
)

// Plan is parsing the HCL code and registering the Job into Nomad
func (co *Client) Validate(job *api.Job) (v *api.JobValidateResponse, err error) {
	v, _, err = co.jobs.Validate(job, nil)
	if err != nil {
		return nil, err
	}

	return v, nil
}
