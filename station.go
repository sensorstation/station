/*

 */
package station

type Station struct {
	ID string // MAC address

	Pubs []Publisher
}

func (s *Station) Start() {

}
