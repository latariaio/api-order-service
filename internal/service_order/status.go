package service_order

type Status string

const (
	StatusOpen             Status = "ABERTA"
	StatusInAnalysis       Status = "EM_ANALISE"
	StatusAwaitingApproval Status = "AGUARDANDO_APROVACAO"
	StatusInProgress       Status = "EM_EXECUCAO"
	StatusCompleted        Status = "CONCLUIDA"
	StatusDelivered        Status = "ENTREGUE"
	StatusCancelled        Status = "CANCELADA"
)

var validTransitions = map[Status][]Status{
	StatusOpen:             {StatusInAnalysis},
	StatusInAnalysis:       {StatusAwaitingApproval, StatusInProgress},
	StatusAwaitingApproval: {StatusInProgress, StatusCancelled},
	StatusInProgress:       {StatusCompleted, StatusCancelled},
	StatusCompleted:        {StatusDelivered},
	StatusDelivered:        {},
	StatusCancelled:        {},
}

func (s Status) CanTransitionTo(next Status) bool {
	allowed, ok := validTransitions[s]
	if !ok {
		return false
	}
	for _, candidate := range allowed {
		if candidate == next {
			return true
		}
	}
	return false
}

func (s Status) IsValid() bool {
	_, ok := validTransitions[s]
	return ok
}

func (s Status) IsClosed() bool {
	return s == StatusCompleted || s == StatusDelivered || s == StatusCancelled
}
