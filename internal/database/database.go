package database

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

type computeLayer interface {
	// парсит и анализирует запрос, если все ок: возвращает set.get или del и аргументы; или ошибку
	HandleQuery(ctx context.Context, input string) (string, []string, error)
}

type storageLayer interface {
	Get(ctx context.Context, key string) (string, bool)
	Set(ctx context.Context, key, value string) error
	Del(ctx context.Context, key string) error
}

type Database struct {
	computeLayer computeLayer
	storageLayer storageLayer
	logger       *zap.Logger
	m            map[string]string
}

func NewDatabase(computeLayer computeLayer, storageLayer storageLayer, logger *zap.Logger) (*Database, error) {
	return &Database{
		computeLayer: computeLayer,
		storageLayer: storageLayer,
		logger:       logger,
		m:            make(map[string]string),
	}, nil
}

func (d *Database) HandleQuery(ctx context.Context, input string) (string, error) {
	command, args, err := d.computeLayer.HandleQuery(ctx, input)
	if err != nil {
		return "", err
	}
	d.logger.Info(fmt.Sprintf("called command is: %s with args: %s", command, args))

	commandSelector := map[string]func(context.Context, []string) (string, error){
		"GET": d.DoGet,
		"SET": d.DoSet,
		"DEL": d.DoDel,
	}

	str, err := commandSelector[command](ctx, args)
	if err != nil {
		return "", err
	}

	return str, nil
}

func (d *Database) DoGet(ctx context.Context, args []string) (string, error) {
	v, ok := d.storageLayer.Get(ctx, args[0])
	if !ok {
		return "", errors.New("key not found")
	}
	return v, nil
}

func (d *Database) DoSet(ctx context.Context, args []string) (string, error) {
	err := d.storageLayer.Set(ctx, args[0], args[1])
	if err != nil {
		return "", err
	}
	return "ok", nil
}

func (d *Database) DoDel(ctx context.Context, args []string) (string, error) {
	err := d.storageLayer.Del(ctx, args[0])
	if err != nil {
		return "", err
	}
	return "ok", nil
}
