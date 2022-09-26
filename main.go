package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type (
	Incidents struct {
		XMLName  xml.Name   `xml:"incidents" json:"incidents,omitempty"`
		Text     string     `xml:",chardata" json:"text,omitempty"`
		Incident []Incident `xml:"incident" json:"incident,omitempty"`
	}
	Incident struct {
		Text string `xml:",chardata" json:"text,omitempty"`
		ID   struct {
			Text string `xml:",chardata" json:"text,omitempty"`
			Type string `xml:"type,attr" json:"type,omitempty"`
		} `xml:"id" json:"id,omitempty"`
		Priority struct {
			Text          string `xml:",chardata" json:"text,omitempty"`
			PriorityLevel struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Type string `xml:"type,attr" json:"type,omitempty"`
			} `xml:"priority-level" json:"priority-level,omitempty"`
			Description string `xml:"description"`
			Hours       struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Type string `xml:"type,attr" json:"type,omitempty"`
			} `xml:"hours" json:"hours,omitempty"`
			CreatedAt struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Type string `xml:"type,attr" json:"type,omitempty"`
			} `xml:"created-at" json:"created-at,omitempty"`
			UpdatedAt struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Type string `xml:"type,attr" json:"type,omitempty"`
			} `xml:"updated-at" json:"updated-at,omitempty"`
			EscalateAt struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Type string `xml:"type,attr" json:"type,omitempty"`
			} `xml:"escalate-at" json:"escalate-at,omitempty"`
			CalendarProfile struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Type string `xml:"type,attr" json:"type,omitempty"`
			} `xml:"calendar-profile" json:"calendar-profile,omitempty"`
			BusinessDaysOnly struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Type string `xml:"type,attr" json:"type,omitempty"`
			} `xml:"business-days-only" json:"business-days-only,omitempty"`
			Level struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Type string `xml:"type,attr" json:"type,omitempty"`
			} `xml:"level" json:"level,omitempty"`
			DeletedAt struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Type string `xml:"type,attr" json:"type,omitempty"`
			} `xml:"deleted-at" json:"deleted-at,omitempty"`
		} `xml:"priority" json:"priority,omitempty"`
		Type struct {
			Text      string `xml:",chardata" json:"text,omitempty"`
			Name      string `xml:"name"`
			CreatedAt struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Type string `xml:"type,attr" json:"type,omitempty"`
			} `xml:"created-at" json:"created-at,omitempty"`
			UpdatedAt struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Type string `xml:"type,attr" json:"type,omitempty"`
			} `xml:"updated-at" json:"updated-at,omitempty"`
			PortalName   string `xml:"portal-name"`
			ShowInPortal struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Type string `xml:"type,attr" json:"type,omitempty"`
			} `xml:"show-in-portal" json:"show-in-portal,omitempty"`
		} `xml:"type" json:"type,omitempty"`
		Service  string `xml:"service"`
		Category struct {
			Text      string `xml:",chardata" json:"text,omitempty"`
			Name      string `xml:"name"`
			CreatedAt struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Type string `xml:"type,attr" json:"type,omitempty"`
			} `xml:"created-at" json:"created-at,omitempty"`
			UpdatedAt struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Type string `xml:"type,attr" json:"type,omitempty"`
			} `xml:"updated-at" json:"updated-at,omitempty"`
			Children struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Type string `xml:"type,attr" json:"type,omitempty"`
			} `xml:"children" json:"children,omitempty"`
			Parent struct {
				Text     string `xml:",chardata" json:"text,omitempty"`
				Name     string `xml:"name"`
				Children struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					Type string `xml:"type,attr" json:"type,omitempty"`
				} `xml:"children" json:"children,omitempty"`
			} `xml:"parent" json:"parent,omitempty"`
		} `xml:"category" json:"category,omitempty"`
		Title string `xml:"title"`
		Owner struct {
			Text      string `xml:",chardata" json:"text,omitempty"`
			FirstName string `xml:"first-name"`
			LastName  string `xml:"last-name"`
			FullName  string `xml:"full-name"`
		} `xml:"owner" json:"owner,omitempty"`
		AssignedTo struct {
			Text      string `xml:",chardata" json:"text,omitempty"`
			FirstName string `xml:"first-name"`
			LastName  string `xml:"last-name"`
			FullName  string `xml:"full-name"`
		} `xml:"assigned-to" json:"assigned-to,omitempty"`
		Customer struct {
			Text      string `xml:",chardata" json:"text,omitempty"`
			FirstName string `xml:"first-name"`
			LastName  string `xml:"last-name"`
		} `xml:"customer" json:"customer,omitempty"`
		Company struct {
			Text string `xml:",chardata" json:"text,omitempty"`
			Name string `xml:"name"`
		} `xml:"company" json:"company,omitempty"`
		RaisedAt         string `xml:"raised-at"`
		OccuredAt        string `xml:"occured-at"`
		DueDate          string `xml:"due-date"`
		ResolvedAt       string `xml:"resolved-at"`
		ClosedAt         string `xml:"closed-at"`
		CustomerViewable struct {
			Text string `xml:",chardata" json:"text,omitempty"`
			Type string `xml:"type,attr" json:"type,omitempty"`
		} `xml:"customer-viewable" json:"customer-viewable,omitempty"`
		ClosureReason  string `xml:"closure-reason"`
		ClosureComment string `xml:"closure-comment"`
		LastUpdatedAt  string `xml:"last-updated-at"`
		LastUpdatedBy  string `xml:"last-updated-by"`
		Symptom        struct {
			Text      string `xml:",chardata" json:"text,omitempty"`
			Note      string `xml:"note"`
			User      string `xml:"user"`
			UpdatedAt string `xml:"updated-at"`
		} `xml:"symptom" json:"symptom,omitempty"`
		Discussion struct {
			Text       string `xml:",chardata" json:"text,omitempty"`
			Type       string `xml:"type,attr" json:"type,omitempty"`
			Discussion []struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Note string `xml:"note" json:"note"`
				User struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					Nil  string `xml:"nil,attr" json:"nil,omitempty"`
				} `xml:"user" json:"user,omitempty"`
				UpdatedAt string `xml:"updated-at" json:"updated_at"`
			} `xml:"discussion" json:"discussion,omitempty"`
		} `xml:"discussion" json:"discussion,omitempty"`
		Resolution struct {
			Text      string `xml:",chardata" json:"text,omitempty"`
			Note      string `xml:"note"`
			User      string `xml:"user"`
			UpdatedAt string `xml:"updated-at"`
		} `xml:"resolution" json:"resolution,omitempty"`
		Incidents        string `xml:"incidents"`
		Problems         string `xml:"problems"`
		Changes          string `xml:"changes"`
		Releases         string `xml:"releases"`
		WatchlistedUsers string `xml:"watchlisted-users"`
		ArchivedAt       string `xml:"archived-at"`
		ResolvedByUserID struct {
			Text string `xml:",chardata" json:"text,omitempty"`
			Type string `xml:"type,attr" json:"type,omitempty"`
		} `xml:"resolved-by-user-id" json:"resolved-by-user-id,omitempty"`
		SlaHoursConsumed struct {
			Text string `xml:",chardata" json:"text,omitempty"`
			Type string `xml:"type,attr" json:"type,omitempty"`
		} `xml:"sla-hours-consumed" json:"sla-hours-consumed,omitempty"`
		SlaHoursStopped struct {
			Text string `xml:",chardata" json:"text,omitempty"`
			Type string `xml:"type,attr" json:"type,omitempty"`
		} `xml:"sla-hours-stopped" json:"sla-hours-stopped,omitempty"`
		Source               string `xml:"source"`
		TimeZone             string `xml:"time-zone"`
		ClockStartedAt       string `xml:"clock-started-at"`
		ClockStoppedAt       string `xml:"clock-stopped-at"`
		ClosedByExternalName string `xml:"closed-by-external-name"`
		ClosedByExternalUuid string `xml:"closed-by-external-uuid"`
		ClosedByUserID       struct {
			Text string `xml:",chardata" json:"text,omitempty"`
			Type string `xml:"type,attr" json:"type,omitempty"`
		} `xml:"closed-by-user-id" json:"closed-by-user-id,omitempty"`
	}
	Discussion struct {
		Text string `json:"text"`
		Note string `json:"note"`
		User struct {
			Text string `json:"text"`
			Nil  string `json:"nil"`
		} `json:"user,omitempty"`
		UpdatedAt string `json:"updated_at"`
	}
)

func main() {
	xmlFile, err := os.Open("incidents.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened incidents.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		log.Fatal(err)
	}
	var incidents Incidents
	if err := xml.Unmarshal(byteValue, &incidents); err != nil {
		log.Fatal(err)
	}

	ticketsFile, err := os.Create("Tickets_01.csv")
	checkError("Cannot create file", err)
	defer ticketsFile.Close()

	writerTicketsFile := csv.NewWriter(ticketsFile)
	defer writerTicketsFile.Flush()

	if err := writerTicketsFile.Write([]string{
		"TicketExtId",
		"ContactExtId",
		"Subject",
		"Description",
		"Email",
		"Due Date",
		"Created Time",
		"Modified Time",
		"Closed Time",
		"Status",
		"Priority",
		"Classification",
		"Owner Email",
		"OwnerExtId",
		"Phone",
		"Category",
		"Sub Category",
		"Resolution",
		"Channel",
	}); err != nil {
		log.Fatal(err)
	}

	threadsFile, err := os.Create("Threads_01.csv")
	checkError("Cannot create file", err)
	defer threadsFile.Close()

	writerThreadsFile := csv.NewWriter(threadsFile)
	defer writerThreadsFile.Flush()
	if err := writerThreadsFile.Write([]string{
		"ThreadExtId",
		"TicketExtId",
		"Content",
		"SendDateTime",
		"ReceivedTime",
		"Sender Email",
		"Direction",
		"isPrivate",
		"FromEmailAddress",
		"To",
		"Cc",
		"Bcc",
		"Channel",
	}); err != nil {
		log.Fatal(err)
	}

	for _, incident := range incidents.Incident {
		if err := writerTicketsFile.Write([]string{
			incident.ID.Text,
			fmt.Sprintf("%s %s", incident.Customer.FirstName, incident.Customer.LastName),
			incident.Title,
			incident.Symptom.Note,
			"?",
			"",
			incident.DueDate,
			incident.LastUpdatedAt,
			incident.ClosedAt,
			"Closed",
			"",
			"",
			incident.AssignedTo.FullName,
			incident.AssignedTo.FullName,
			fmt.Sprintf("%s %s", incident.Customer.FirstName, incident.Customer.LastName),
			incident.Category.Parent.Name,
			incident.Category.Name,
			incident.Resolution.Note,
			incident.Source,
		}); err != nil {
			log.Fatal(err)
		}
		for idx, discussion := range incident.Discussion.Discussion {
			if err := writerThreadsFile.Write([]string{
				fmt.Sprintf("ticket_%s_%d", incident.ID.Text, idx+1),
				incident.ID.Text,
				discussion.Note,
				incident.OccuredAt,
				incident.DueDate,
				"",
				"",
				incident.CustomerViewable.Text,
				"",
				"",
				"",
				"",
				incident.Source,
			}); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
