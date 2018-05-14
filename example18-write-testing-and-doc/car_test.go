package car

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Simple testing what different between Fatal and Error
func TestNew(t *testing.T) {
	c, err := New("", 100)
	if err != nil {
		t.Fatal("got errors:", err)
	}

	if c == nil {
		t.Error("car shoud be nil")
	}
}

// Simple testing with testify tool
func TestNewWithAssert(t *testing.T) {
	c, err := New("", 100)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Nil(t, c)

	c, err = New("foo", 100)
	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, "foo", c.Name)
}

func TestCar_SetName(t *testing.T) {
	type fields struct {
		Name  string
		Price float32
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "no input name",
			fields: fields{
				Name:  "foo",
				Price: 100,
			},
			args: args{
				name: "",
			},
			want: "foo",
		},
		{
			name: "input name",
			fields: fields{
				Name:  "foo",
				Price: 100,
			},
			args: args{
				name: "bar",
			},
			want: "bar",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := &Car{
				Name:  tt.fields.Name,
				Price: tt.fields.Price,
			}
			if got := c.SetName(tt.args.name); got != tt.want {
				t.Errorf("Car.SetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
