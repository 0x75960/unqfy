package unqfy

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
)

//Listup file in directory
func Listup(root string) (files []string, err error) {

	err = filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {

			if info.IsDir() {
				return nil
			}

			abs, err := filepath.Abs(path)
			if err != nil {
				return err
			}

			files = append(files, abs)

			return nil
		})

	if err != nil {
		return files, err
	}

	return
}

// ListupInDirs specified
func ListupInDirs(dirs []string) (all []string, err error) {

	for _, dir := range dirs {

		files, err := Listup(dir)
		if err != nil {
			return all, err
		}

		for _, file := range files {
			all = append(all, file)
		}

	}

	return

}

// Uniqify files
func Uniqify(files []string) (uniqified []string, err error) {

	sha256ToPath := make(map[string][]string)

	for _, file := range files {

		sum, err := sha256sum(file)
		if err != nil {
			return uniqified, err
		}

		// sha256でまとめる
		sha256ToPath[sum] = append(sha256ToPath[sum], file)
	}

	for _, value := range sha256ToPath {
		uniqified = append(uniqified, value[0])
	}

	return
}

// sha256sum of file
func sha256sum(file string) (sum string, err error) {

	sha256er := sha256.New()

	f, err := os.Open(file)
	if err != nil {
		return sum, err
	}
	defer f.Close()

	_, err = io.Copy(sha256er, f)
	if err != nil {
		return sum, err
	}

	return hex.EncodeToString(sha256er.Sum(nil)), nil
}

// Copy unique files in directory.
func Copy(dst, src string) (err error) {

	// フォルダの中のファイルを列挙
	all, err := Listup(src)
	if err != nil {
		return err
	}

	// 見つけたファイルをユニークに
	uniqified, err := Uniqify(all)
	if err != nil {
		return err
	}

	// 出力先ディレクトリが存在しない場合作成
	_, err = os.Stat(dst)
	if err != nil {

		err = os.MkdirAll(dst, os.ModePerm)
		if err != nil {
			return err
		}

		log.Printf("%s not found. created.", dst)

	}

	for _, file := range uniqified {

		sum, err := sha256sum(file)
		if err != nil {
			return err
		}

		// 出力先ディレクトリにSHA256のファイル名でコピー
		dstPath := path.Join(dst, sum)
		err = filecopy(dstPath, file)
		if err != nil {
			return err
		}

	}

	return

}

// filecopy to dst from src
func filecopy(dst, src string) (err error) {

	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer d.Close()

	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()

	_, err = io.Copy(d, s)

	return
}
