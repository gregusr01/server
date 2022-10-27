// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package token

// GetRepo gets a repo by ID from the database.
func (e *engine) GetInvalidToken(t string) error {
	e.logger.Tracef("getting token hash from the database")

	// variable to store query results
	var ts string

	// send query to the database and store result in variable
	err := e.client.
		Table("invalid_tokens").
		Where("token_hash = ?", t).
		Take(ts).
		Error
	if err != nil {
		return err
	}

	e.logger.Tracef("what we got back: %s", ts)

	// if we got something
	// return errors.New("token hash found inside invalidation database")

	return nil
}
