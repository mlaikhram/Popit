import React from 'react';
import { ScrollMenu } from "react-horizontal-scrolling-menu";
import { LeftArrow, RightArrow } from "./Arrows";


export function PageScrollMenu(props) {
    return (
      <>
        <h2>{props.name}</h2>
        <ScrollMenu LeftArrow={LeftArrow} RightArrow={RightArrow} style={{ backgroundColor: '#131313' }}>{props.items}</ScrollMenu>
      </>
    );
  }
