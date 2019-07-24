package model

func DBUpdatePersonalInfo(newUserInfo RegisterJSON) error {
	stmt, err := DB.Prepare("UPDATE users SET nickname=?, mobile=?, email=? WHERE username=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(newUserInfo.Nickname, newUserInfo.Mobile, newUserInfo.Email, newUserInfo.Username)
	if err != nil {
		return err
	}

	return nil
}

func DBUpdatePassword(username, password string) error {
	stmt, err := DB.Prepare("UPDATE users SET password=? WHERE username=?")
	if err != nil {
		return err
	}

	hashedPassword, err := encryptPassword(password)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(hashedPassword, username)
	if err != nil {
		return err
	}
	return nil
}
