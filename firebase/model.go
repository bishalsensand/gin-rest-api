package firebase

type Entry struct {
	Id                *string `json:"id" firestore:"id,omitempty"`
	AccountName       *string `json:"account_name" binding:"required" firestore:"account_name,omitempty"`
	ContactName       *string `json:"contact_name" binding:"required" firestore:"contact_name,omitempty"`
	DealSize          *int    `json:"deal_size" binding:"required" firestore:"deal_size,omitempty"`
	SalesRep          *string `json:"sales_rep" binding:"required" firestore:"sales_rep,omitempty"`
	RegisteredWebinar *bool   `json:"registered_webinar" binding:"required" firestore:"registered_webinar,omitempty"`
	SalesCall         *bool   `json:"sales_call" binding:"required" firestore:"sales_call,omitempty"`
	DealStatus        *bool   `json:"deal_status" binding:"required" firestore:"deal_status,omitempty"`
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
	return deleteFirebaseEntryById(id)
}

func (e Entry) UpdateById(entryId string) (bool, error) {
	return updateFirebaseEntryById(entryId, e)
}
