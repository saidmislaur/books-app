package author

import (
	"fmt"
)

func (am *AuthorManager) Delete(id int) error {
	res, err := am.DB.Exec(`DELETE FROM authors WHERE id = $1`, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("автор с id %d не найден", id)
	}

	return nil
}
