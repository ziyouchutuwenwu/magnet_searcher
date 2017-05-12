package link

func MagnetLink(infoHash string) string{

	magnetPrefix := "magnet:?xt=urn:btih:"

	return magnetPrefix + infoHash;
}