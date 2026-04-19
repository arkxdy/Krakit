package messaging

const (
	TopicAttemptStarted            = "attempt.started"             //sync
	TopicAttemptSubmitted          = "attempt.submitted"           //worker pool
	TopicAttemptEvaluated          = "attempt.evaluated"           //worker pool
	TopicAnswerSaved               = "answer.saved"                //sync
	TopicQuestionBulkUploaded      = "question.bulk.uploaded"      //worker pool
	TopicQuestionUpdated           = "question.updated"            //sync
	TopicExamPublished             = "exam.published"              //kafka
	TopicAnalyticsAttemptProcessed = "analytics.attempt.processed" //kafka
)
