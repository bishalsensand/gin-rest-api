package firebase

type Entry struct {
	Id                string `json:"id" firestore:"id,omitempty"`
	AccountName       string `json:"account_name" binding:"required"`
	ContactName       string `json:"contact_name" binding:"required"`
	DealSize          *int   `json:"deal_size" binding:"required"`
	SalesRep          string `json:"sales_rep" binding:"required"`
	RegisteredWebinar *bool  `json:"registered_webinar" binding:"required"`
	SalesCall         *bool  `json:"sales_call" binding:"required"`
	DealStatus        *bool  `json:"deal_status" binding:"required"`
}

func (e Entry) CreateNew() (string, error) {
	return postEntryToFirebase(e)
}

func GetEntryById(id string) (*Entry, error) {
	return getEntryFromFirebase(id)
}

func GetEntries(count int) ([]Entry, error) {
	return getAllEntriesFromFirebase(count)
}

func DeleteEntryById(id string) (bool, error) {
	return deleteEntryById(id)
}
