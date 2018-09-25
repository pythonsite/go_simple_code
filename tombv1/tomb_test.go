package tombv1

import (
	"reflect"
	"testing"
	"gopkg.in/tomb.v1"
	"errors"
)

func TestNewTomb(t *testing.T) {
	tb := &tomb.Tomb{}
	testState(t,tb,false,false,tomb.ErrStillAlive)

	tb.Done()
	testState(t,tb,true,true,nil)
}

func TestKill(t *testing.T) {
	tb := &tomb.Tomb{}
	tb.Kill(nil)
	testState(t,tb,true,false,nil)

	err := errors.New("some error")
	tb.Kill(err)
	testState(t, tb, true, false, err)

	tb.Kill(errors.New("ignore me"))
	testState(t, tb,true,false,err)

	tb.Done()
	testState(t,tb,true,true,err)
}



func testState(t *testing.T, tb *tomb.Tomb, wantDying, wantDead bool, wantErr error) {
	select {
	case <- tb.Dying():
		if !wantDying {
			t.Error("<-Dying:should block")
		}
	default:
		if wantDying {
			t.Error("<-DYing: should not block")
		}
	}

	seemsDead := false
	select {
	case <- tb.Dead():
		if !wantDead {
			t.Error("<-Dead: should block")
		}
		seemsDead = true
	default:
		if wantDead {
			t.Error("<-Dead: should not block")
		}
	}


	if err := tb.Err();err!= wantErr {
		t.Errorf("Err: want %#v, got %#v", wantErr, err)
	}

	if wantDead && seemsDead {
		waitErr := tb.Wait()
		switch {
		case waitErr == tomb.ErrStillAlive:
			t.Errorf("Wait should not return ErrStillAlive")
		case !reflect.DeepEqual(waitErr, wantErr):
			t.Errorf("Want: want %#v, got %#v", wantErr, waitErr)
		}
	}
}
