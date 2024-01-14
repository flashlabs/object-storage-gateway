package storage_test

import (
	"errors"
	"testing"

	"github.com/spacelift-io/homework-object-storage/internal/registry"
	"github.com/spacelift-io/homework-object-storage/internal/service/storage"
	"github.com/spacelift-io/homework-object-storage/internal/structs"
	"github.com/spacelift-io/homework-object-storage/pkg"
)

var (
	oneShard = map[uint8]structs.Storage{
		uint8(0): {Client: nil},
	}
	twoShards = map[uint8]structs.Storage{
		uint8(0): {Client: nil},
		uint8(1): {Client: nil},
	}
	threeShards = map[uint8]structs.Storage{
		uint8(0): {Client: nil},
		uint8(1): {Client: nil},
		uint8(2): {Client: nil},
	}
	numID1 = "1234567890"
	numID2 = "12345678901"
	alpha1 = "lollipop"
	alpha2 = "gingerbread"
	alnum1 = "cake123"
	alnum2 = "cupcake1234567890soufflé"
)

func TestShardByID(t *testing.T) {
	type args struct {
		id     string
		shards map[uint8]structs.Storage
	}
	tests := []struct {
		name    string
		args    args
		want    uint8
		wantErr bool
		err     error
	}{
		{
			name:    "expect error bc of empty shards",
			args:    args{id: numID1},
			want:    uint8(0),
			wantErr: true,
			err:     pkg.ErrNoShardsAvailable,
		},
		{
			name: "test numID1 on 1 shard",
			args: args{
				id:     numID1,
				shards: oneShard,
			},
			want:    uint8(0),
			wantErr: false,
		},
		{
			name: "test numID2 on 1 shard",
			args: args{
				id:     numID2,
				shards: oneShard,
			},
			want:    uint8(0),
			wantErr: false,
		},
		{
			name: "test numID1 on 2 shards",
			args: args{
				id:     numID1,
				shards: twoShards,
			},
			want:    uint8(0),
			wantErr: false,
		},
		{
			name: "test numID2 on 2 shards",
			args: args{
				id:     numID2,
				shards: twoShards,
			},
			want:    uint8(1),
			wantErr: false,
		},
		{
			name: "test numID1 on 3 shards",
			args: args{
				id:     numID1,
				shards: threeShards,
			},
			want:    uint8(0),
			wantErr: false,
		},
		{
			name: "test numID2 on 3 shards",
			args: args{
				id:     numID2,
				shards: threeShards,
			},
			want:    uint8(1),
			wantErr: false,
		},
		{
			name: "test alpha1 on 3 shards",
			args: args{
				id:     alpha1,
				shards: threeShards,
			},
			want:    uint8(0),
			wantErr: false,
		},
		{
			name: "test alpha2 on 3 shards",
			args: args{
				id:     alpha2,
				shards: threeShards,
			},
			want:    uint8(2),
			wantErr: false,
		},
		{
			name: "test alnum1 on 3 shards",
			args: args{
				id:     alnum1,
				shards: threeShards,
			},
			want:    uint8(2),
			wantErr: false,
		},
		{
			name: "test alnum2 on 3 shards",
			args: args{
				id:     alnum2,
				shards: threeShards,
			},
			want:    uint8(0),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registry.Shards = tt.args.shards

			got, err := storage.ShardByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShardByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("ShardByID() got = %v, want %v", got, tt.want)
			}

			if tt.wantErr && tt.err != nil {
				errors.Is(err, tt.err)
			}
		})
	}
}
