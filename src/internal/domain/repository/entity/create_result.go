package entity

func NewCreateResult(id ID) CreateResult {
	return CreateResult{
		id: id,
	}
}

type CreateResult struct {
	id ID
}

func (c *CreateResult) GetID() ID {
	return c.id
}
