// Code generated by "stringer -type=AgreementType"; DO NOT EDIT.

package logspec

import "strconv"

const _AgreementType_name = "RoundConcludedPeriodConcludedStepTimeoutRoundStartRoundInterruptedRoundWaitingThresholdReachedBlockAssembledBlockCommittableProposalAssembledProposalBroadcastProposalFrozenProposalAcceptedProposalRejectedBlockRejectedBlockResentBlockPipelinedVoteAttestVoteBroadcastVoteAcceptedVoteRejectedBundleBroadcastBundleAcceptedBundleRejectedRestoredPersistednumAgreementTypes"

var _AgreementType_index = [...]uint16{0, 14, 29, 40, 50, 66, 78, 94, 108, 124, 141, 158, 172, 188, 204, 217, 228, 242, 252, 265, 277, 289, 304, 318, 332, 340, 349, 366}

func (i AgreementType) String() string {
	if i < 0 || i >= AgreementType(len(_AgreementType_index)-1) {
		return "AgreementType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AgreementType_name[_AgreementType_index[i]:_AgreementType_index[i+1]]
}