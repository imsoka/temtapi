package repository

import "temtapi/internal/domain"

type TemtemRespository interface {
    Store(t domain.Temtem) *domain.Temtem
    getById(id string) *domain.Temtem
    getByTempediaNumber(n uint8) *domain.Temtem
    GetByName(n string) *domain.Temtem
}
