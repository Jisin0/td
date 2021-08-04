package updates

import (
	"golang.org/x/xerrors"

	"github.com/gotd/td/tg"
)

func validatePts(pts, ptsCount int) error {
	if pts < 0 {
		return xerrors.Errorf("invalid pts value: %d", pts)
	}

	if ptsCount < 0 {
		return xerrors.Errorf("invalid ptsCount value: %d", ptsCount)
	}

	return nil
}

func validateQts(qts int) error {
	if qts < 0 {
		return xerrors.Errorf("invalid qts value: %d", qts)
	}

	return nil
}

func validateSeq(seq, seqStart int) error {
	if seq < 0 {
		return xerrors.Errorf("invalid seq value: %d", seq)
	}

	if seqStart < 0 {
		return xerrors.Errorf("invalid seqStart value: %d", seq)
	}

	return nil
}

func getDialogPts(dialog tg.DialogClass) (int, error) {
	d, ok := dialog.(*tg.Dialog)
	if !ok {
		return 0, xerrors.Errorf("unexpected dialog type: %T", dialog)
	}

	pts, ok := d.GetPts()
	if !ok {
		return 0, xerrors.New("dialog has no pts field")
	}

	return pts, nil
}
