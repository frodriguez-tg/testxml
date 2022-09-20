package main

import "encoding/xml"

type Incidents struct {
	XMLName  xml.Name `xml:"incidents"`
	Text     string   `xml:",chardata"`
	Incident []struct {
		Text string `xml:",chardata"`
		ID   struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"id"`
		Priority struct {
			Text          string `xml:",chardata"`
			PriorityLevel struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"priority-level"`
			Description string `xml:"description"`
			Hours       struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"hours"`
			CreatedAt struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"created-at"`
			UpdatedAt struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"updated-at"`
			EscalateAt struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"escalate-at"`
			CalendarProfile struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"calendar-profile"`
			BusinessDaysOnly struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"business-days-only"`
			Level struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"level"`
			DeletedAt struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"deleted-at"`
		} `xml:"priority"`
		Type struct {
			Text      string `xml:",chardata"`
			Name      string `xml:"name"`
			CreatedAt struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"created-at"`
			UpdatedAt struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"updated-at"`
			PortalName   string `xml:"portal-name"`
			ShowInPortal struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"show-in-portal"`
		} `xml:"type"`
		Service  string `xml:"service"`
		Category struct {
			Text      string `xml:",chardata"`
			Name      string `xml:"name"`
			CreatedAt struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"created-at"`
			UpdatedAt struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"updated-at"`
			Children struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"children"`
			Parent struct {
				Text     string `xml:",chardata"`
				Name     string `xml:"name"`
				Children struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"children"`
			} `xml:"parent"`
		} `xml:"category"`
		Title string `xml:"title"`
		Owner struct {
			Text      string `xml:",chardata"`
			FirstName string `xml:"first-name"`
			LastName  string `xml:"last-name"`
			FullName  string `xml:"full-name"`
		} `xml:"owner"`
		AssignedTo struct {
			Text      string `xml:",chardata"`
			FirstName string `xml:"first-name"`
			LastName  string `xml:"last-name"`
			FullName  string `xml:"full-name"`
		} `xml:"assigned-to"`
		Customer struct {
			Text      string `xml:",chardata"`
			FirstName string `xml:"first-name"`
			LastName  string `xml:"last-name"`
		} `xml:"customer"`
		Company struct {
			Text string `xml:",chardata"`
			Name string `xml:"name"`
		} `xml:"company"`
		RaisedAt         string `xml:"raised-at"`
		OccuredAt        string `xml:"occured-at"`
		DueDate          string `xml:"due-date"`
		ResolvedAt       string `xml:"resolved-at"`
		ClosedAt         string `xml:"closed-at"`
		CustomerViewable struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"customer-viewable"`
		ClosureReason  string `xml:"closure-reason"`
		ClosureComment string `xml:"closure-comment"`
		LastUpdatedAt  string `xml:"last-updated-at"`
		LastUpdatedBy  string `xml:"last-updated-by"`
		Symptom        struct {
			Text      string `xml:",chardata"`
			Note      string `xml:"note"`
			User      string `xml:"user"`
			UpdatedAt string `xml:"updated-at"`
		} `xml:"symptom"`
		Discussion struct {
			Text       string `xml:",chardata"`
			Type       string `xml:"type,attr"`
			Discussion []struct {
				Text string `xml:",chardata"`
				Note string `xml:"note"`
				User struct {
					Text string `xml:",chardata"`
					Nil  string `xml:"nil,attr"`
				} `xml:"user"`
				UpdatedAt string `xml:"updated-at"`
			} `xml:"discussion"`
		} `xml:"discussion"`
		Resolution struct {
			Text      string `xml:",chardata"`
			Note      string `xml:"note"`
			User      string `xml:"user"`
			UpdatedAt string `xml:"updated-at"`
		} `xml:"resolution"`
		Incidents        string `xml:"incidents"`
		Problems         string `xml:"problems"`
		Changes          string `xml:"changes"`
		Releases         string `xml:"releases"`
		WatchlistedUsers string `xml:"watchlisted-users"`
		ArchivedAt       string `xml:"archived-at"`
		ResolvedByUserID struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"resolved-by-user-id"`
		SlaHoursConsumed struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"sla-hours-consumed"`
		SlaHoursStopped struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"sla-hours-stopped"`
		Source               string `xml:"source"`
		TimeZone             string `xml:"time-zone"`
		ClockStartedAt       string `xml:"clock-started-at"`
		ClockStoppedAt       string `xml:"clock-stopped-at"`
		ClosedByExternalName string `xml:"closed-by-external-name"`
		ClosedByExternalUuid string `xml:"closed-by-external-uuid"`
		ClosedByUserID       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"closed-by-user-id"`
	} `xml:"incident"`
}

func main() {

}
