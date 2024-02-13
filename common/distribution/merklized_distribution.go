package distribution

type MerklizedDistribution struct {
	distribution Distribution
}

func NewMerklizedDistribution(distribution Distribution) *MerklizedDistribution {
	return &MerklizedDistribution{
		distribution: distribution,
	}

}
