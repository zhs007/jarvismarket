package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jarviscore"
	jarvismarketbasedef "github.com/zhs007/jarvismarket/basedef"
	"github.com/zhs007/jarvismarket/jarvismarket"
)

func startServ() {
	fmt.Printf("jarvismarket start...\n")
	fmt.Printf("jarvismarket version is %v \n", jarvismarketbasedef.VERSION)

	cfg, err := jarviscore.LoadConfig("cfg/jarvisnode.yaml")
	if err != nil {
		fmt.Printf("load jarvisnode.yaml fail!\n")

		return
	}

	jarviscore.InitJarvisCore(cfg)
	defer jarviscore.ReleaseJarvisCore()

	market, err := jarvismarket.NewMarket("./cfg/config.yaml")
	if err != nil {
		fmt.Printf("jarvismarket.NewMarket %v", err)

		return
	}

	// pprof
	jarviscore.InitPprof(cfg)

	node, err := jarviscore.NewNode(cfg)
	if err != nil {
		fmt.Printf("jarviscore.NewNode fail! %v \n", err)

		return
	}

	node.SetNodeTypeInfo(jarvismarketbasedef.JARVISNODETYPE, jarvismarketbasedef.VERSION)

	go market.Start(context.Background())
	node.Start(context.Background())

	fmt.Printf("jarvismarket end.\n")
}
