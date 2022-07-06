package handlers

import . "github.com/PSE-SS2022/timefy-backend/internal/models"

type ReportHandler struct {
}

func (reportHandler ReportHandler) ReportUser(userId, reportedUserId, reportReason string) {

}

func (reportHandler ReportHandler) ReportEvent(userId, eventToReportId, reportReason string) {

}

func (reportHandler ReportHandler) createReport(reporter User, reportedObject string, reason string) Report {
	var result Report
	return result
}

func (reportHandler ReportHandler) makeObjectSnapshot(objectId string) string {
	return ""
}
