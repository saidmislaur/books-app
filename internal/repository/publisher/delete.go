package publisher

import (
	"fmt"
)

func (p *PublisherManager) Delete(id int) error {
	res, err := p.DB.Exec(`DELETE FROM publishers WHERE id = $1`, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("издательство с id: %d не найдено", id)
	}

	return nil
}
