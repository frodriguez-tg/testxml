package main

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type (
	Incidents struct {
		XMLName  xml.Name `xml:"incidents" json:"incidents,omitempty"`
		Text     string   `xml:",chardata" json:"text,omitempty"`
		Incident []struct {
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
		} `xml:"incident" json:"incident,omitempty"`
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

	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{
		"id",

		"priority.PriorityLevel",
		"priority.Description",
		"priority.Hours",
		"priority.CreatedAt",
		"priority.UpdatedAt",
		"priority.EscalateAt",
		"priority.CalendarProfile",
		"priority.BusinessDaysOnly",
		"priority.Level",
		"priority.DeletedAt",

		"type",
		"type.Name",
		"type.CreatedAt",
		"type.UpdatedAt",
		"type.PortalName",
		"type.ShowInPortal",

		"service",

		"category",
		"category.Name",
		"category.CreatedAt",
		"category.UpdatedAt",
		"category.Children",
		"category.Parent",
		"category.Parent.Name",
		"category.Parent.Children",

		"title",

		"owner",
		"owner.FirstName",
		"owner.LastName",
		"owner.FullName",

		"assignedTo",
		"assignedTo.FirstName",
		"assignedTo.LastName",
		"assignedTo.FullName",

		"customer",
		"customer.FirstName",
		"customer.LastName",

		"company",
		"company.Name",

		"raisedAt",
		"occuredAt",
		"dueDate",
		"resolvedAt",
		"closedAt",

		"customerViewable",

		"closureReason",
		"closureComment",
		"lastUpdatedAt",
		"lastUpdatedBy",

		"symptom",
		"symptom.Note",
		"symptom.User",
		"symptom.UpdatedAt",

		"discussion",
		"arrayDiscussions",

		"resolution",
		"resolution.Note",
		"resolution.User",
		"resolution.UpdatedAt",

		"incidents",
		"problems",
		"changes",
		"releases",
		"watchlistedUsers",
		"archivedAt",

		"resolvedByUserID",

		"slaHoursConsumed",

		"slaHoursStopped",

		"source",
		"timeZone",
		"clockStartedAt",
		"clockStoppedAt",
		"closedByExternalName",
		"closedByExternalUuid",

		"closedByUserID",
	}); err != nil {
		log.Fatal(err)
	}

	for _, result := range incidents.Incident {

		discussions := make([]Discussion, 0)
		for _, discussion := range result.Discussion.Discussion {
			discussions = append(discussions, Discussion{
				Text: strings.ReplaceAll(strings.TrimSpace(discussion.Text), "\n", ""),
				Note: strings.ReplaceAll(strings.TrimSpace(discussion.Note), "\n", ""),
				User: struct {
					Text string `json:"text"`
					Nil  string `json:"nil"`
				}{
					Text: strings.ReplaceAll(strings.TrimSpace(discussion.User.Text), "\n", ""),
					Nil:  strings.ReplaceAll(strings.TrimSpace(discussion.User.Nil), "\n", ""),
				},
				UpdatedAt: strings.ReplaceAll(strings.TrimSpace(discussion.UpdatedAt), "\n", ""),
			})
		}

		discussionBytes, err := json.Marshal(discussions)
		if err != nil {
			log.Fatal(err)
		}

		err = writer.Write([]string{
			strings.TrimSpace(result.ID.Text),

			strings.TrimSpace(result.Priority.PriorityLevel.Text),
			strings.TrimSpace(result.Priority.Description),
			strings.TrimSpace(result.Priority.Hours.Text),
			strings.TrimSpace(result.Priority.CreatedAt.Text),
			strings.TrimSpace(result.Priority.UpdatedAt.Text),
			strings.TrimSpace(result.Priority.EscalateAt.Text),
			strings.TrimSpace(result.Priority.CalendarProfile.Text),
			strings.TrimSpace(result.Priority.BusinessDaysOnly.Text),
			strings.TrimSpace(result.Priority.Level.Text),
			strings.TrimSpace(result.Priority.DeletedAt.Text),

			strings.TrimSpace(result.Type.Text),
			strings.TrimSpace(result.Type.Name),
			strings.TrimSpace(result.Type.CreatedAt.Text),
			strings.TrimSpace(result.Type.UpdatedAt.Text),
			strings.TrimSpace(result.Type.PortalName),
			strings.TrimSpace(result.Type.ShowInPortal.Text),

			strings.TrimSpace(result.Service),

			strings.TrimSpace(result.Category.Text),
			strings.TrimSpace(result.Category.Name),
			strings.TrimSpace(result.Category.CreatedAt.Text),
			strings.TrimSpace(result.Category.UpdatedAt.Text),
			strings.TrimSpace(result.Category.Children.Text),
			strings.TrimSpace(result.Category.Parent.Text),
			strings.TrimSpace(result.Category.Parent.Name),
			strings.TrimSpace(result.Category.Parent.Children.Text),

			strings.TrimSpace(result.Title),

			strings.TrimSpace(result.Owner.Text),
			strings.TrimSpace(result.Owner.FirstName),
			strings.TrimSpace(result.Owner.LastName),
			strings.TrimSpace(result.Owner.FullName),

			strings.TrimSpace(result.AssignedTo.Text),
			strings.TrimSpace(result.AssignedTo.FirstName),
			strings.TrimSpace(result.AssignedTo.LastName),
			strings.TrimSpace(result.AssignedTo.FullName),

			strings.TrimSpace(result.Customer.Text),
			strings.TrimSpace(result.Customer.FirstName),
			strings.TrimSpace(result.Customer.LastName),

			strings.TrimSpace(result.Company.Text),
			strings.TrimSpace(result.Company.Name),

			strings.TrimSpace(result.RaisedAt),
			strings.TrimSpace(result.OccuredAt),
			strings.TrimSpace(result.DueDate),
			strings.TrimSpace(result.ResolvedAt),
			strings.TrimSpace(result.ClosedAt),

			strings.TrimSpace(result.CustomerViewable.Text),

			strings.TrimSpace(result.ClosureReason),
			strings.TrimSpace(result.ClosureComment),
			strings.TrimSpace(result.LastUpdatedAt),
			strings.TrimSpace(result.LastUpdatedBy),

			strings.TrimSpace(result.Symptom.Text),
			strings.TrimSpace(result.Symptom.Note),
			strings.TrimSpace(result.Symptom.User),
			strings.TrimSpace(result.Symptom.UpdatedAt),

			strings.TrimSpace(result.Discussion.Text),
			strings.TrimSpace(string(discussionBytes)),

			strings.TrimSpace(result.Resolution.Text),
			strings.TrimSpace(result.Resolution.Note),
			strings.TrimSpace(result.Resolution.User),
			strings.TrimSpace(result.Resolution.UpdatedAt),

			strings.TrimSpace(result.Incidents),
			strings.TrimSpace(result.Problems),
			strings.TrimSpace(result.Changes),
			strings.TrimSpace(result.Releases),
			strings.TrimSpace(result.WatchlistedUsers),
			strings.TrimSpace(result.ArchivedAt),

			strings.TrimSpace(result.ResolvedByUserID.Text),

			strings.TrimSpace(result.SlaHoursConsumed.Text),

			strings.TrimSpace(result.SlaHoursStopped.Text),

			strings.TrimSpace(result.Source),
			strings.TrimSpace(result.TimeZone),
			strings.TrimSpace(result.ClockStartedAt),
			strings.TrimSpace(result.ClockStoppedAt),
			strings.TrimSpace(result.ClosedByExternalName),
			strings.TrimSpace(result.ClosedByExternalUuid),

			strings.TrimSpace(result.ClosedByUserID.Text),
		})
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
