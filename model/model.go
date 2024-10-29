package model

import "time"

type NotionRespBody struct {
	Object  string `json:"object"`
	Results []struct {
		Object         string    `json:"object"`
		ID             string    `json:"id"`
		CreatedTime    time.Time `json:"created_time"`
		LastEditedTime time.Time `json:"last_edited_time"`
		CreatedBy      struct {
			Object string `json:"object"`
			ID     string `json:"id"`
		} `json:"created_by"`
		LastEditedBy struct {
			Object string `json:"object"`
			ID     string `json:"id"`
		} `json:"last_edited_by"`
		Cover  any `json:"cover"`
		Icon   any `json:"icon"`
		Parent struct {
			Type       string `json:"type"`
			DatabaseID string `json:"database_id"`
		} `json:"parent"`
		Archived   bool `json:"archived"`
		InTrash    bool `json:"in_trash"`
		Properties struct {
			JadwalKeberangkatan struct {
				ID   string `json:"id"`
				Type string `json:"type"`
				Date any    `json:"date"`
			} `json:"Jadwal Keberangkatan"`
			Lokasi struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				RichText []struct {
					Type string `json:"type"`
					Text struct {
						Content string `json:"content"`
						Link    any    `json:"link"`
					} `json:"text"`
					Annotations struct {
						Bold          bool   `json:"bold"`
						Italic        bool   `json:"italic"`
						Strikethrough bool   `json:"strikethrough"`
						Underline     bool   `json:"underline"`
						Code          bool   `json:"code"`
						Color         string `json:"color"`
					} `json:"annotations"`
					PlainText string `json:"plain_text"`
					Href      any    `json:"href"`
				} `json:"rich_text"`
			} `json:"Lokasi"`
			NoWA08Xxxxx struct {
				ID          string `json:"id"`
				Type        string `json:"type"`
				PhoneNumber string `json:"phone_number"`
			} `json:"No WA (08xxxxx)"`
			BuktiTransfer struct {
				ID    string `json:"id"`
				Type  string `json:"type"`
				Files []any  `json:"files"`
			} `json:"Bukti Transfer"`
			FakturTagihan struct {
				ID    string `json:"id"`
				Type  string `json:"type"`
				Files []any  `json:"files"`
			} `json:"Faktur Tagihan"`
			KuitansiPembayaran struct {
				ID    string `json:"id"`
				Type  string `json:"type"`
				Files []any  `json:"files"`
			} `json:"Kuitansi Pembayaran"`
			KelengkapanDokumen struct {
				ID          string `json:"id"`
				Type        string `json:"type"`
				MultiSelect []any  `json:"multi_select"`
			} `json:"Kelengkapan dokumen"`
			Status struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Status struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Color string `json:"color"`
				} `json:"status"`
			} `json:"Status"`
			Produk struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Select any    `json:"select"`
			} `json:"Produk"`
			AccountOwner struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				People []struct {
					Object    string `json:"object"`
					ID        string `json:"id"`
					Name      string `json:"name"`
					AvatarURL string `json:"avatar_url"`
					Type      string `json:"type"`
					Person    struct {
						Email string `json:"email"`
					} `json:"person"`
				} `json:"people"`
			} `json:"Account Owner"`
			LampiranDokumen struct {
				ID    string `json:"id"`
				Type  string `json:"type"`
				Files []any  `json:"files"`
			} `json:"Lampiran Dokumen"`
			LeadsCreated struct {
				ID   string `json:"id"`
				Type string `json:"type"`
				Date struct {
					Start    time.Time `json:"start"`
					End      any       `json:"end"`
					TimeZone any       `json:"time_zone"`
				} `json:"date"`
			} `json:"Leads Created"`
			ClosingRate struct {
				ID      string `json:"id"`
				Type    string `json:"type"`
				Formula struct {
					Type   string `json:"type"`
					String any    `json:"string"`
				} `json:"formula"`
			} `json:"Closing Rate"`
			WAGroup struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Select struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Color string `json:"color"`
				} `json:"select"`
			} `json:"WA Group"`
			LastEditedTime struct {
				ID             string    `json:"id"`
				Type           string    `json:"type"`
				LastEditedTime time.Time `json:"last_edited_time"`
			} `json:"Last edited time"`
			TotalTagihan struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Number any    `json:"number"`
			} `json:"Total Tagihan"`
			JumlahAnggotaRombongan struct {
				ID     string `json:"id"`
				Type   string `json:"type"`
				Number any    `json:"number"`
			} `json:"Jumlah Anggota Rombongan"`
			Name struct {
				ID    string `json:"id"`
				Type  string `json:"type"`
				Title []struct {
					Type string `json:"type"`
					Text struct {
						Content string `json:"content"`
						Link    any    `json:"link"`
					} `json:"text"`
					Annotations struct {
						Bold          bool   `json:"bold"`
						Italic        bool   `json:"italic"`
						Strikethrough bool   `json:"strikethrough"`
						Underline     bool   `json:"underline"`
						Code          bool   `json:"code"`
						Color         string `json:"color"`
					} `json:"annotations"`
					PlainText string `json:"plain_text"`
					Href      any    `json:"href"`
				} `json:"title"`
			} `json:"Name"`
		} `json:"properties"`
		URL       string `json:"url"`
		PublicURL any    `json:"public_url"`
	} `json:"results"`
	NextCursor     any    `json:"next_cursor"`
	HasMore        bool   `json:"has_more"`
	Type           string `json:"type"`
	PageOrDatabase struct {
	} `json:"page_or_database"`
	DeveloperSurvey string `json:"developer_survey"`
	RequestID       string `json:"request_id"`
}
