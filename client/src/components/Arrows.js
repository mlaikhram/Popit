import React from "react";

import { VisibilityContext } from "react-horizontal-scrolling-menu";

function Arrow({
  children,
  disabled,
  onClick,
  marginLeft,
  marginRight
}) {
  return (
    <button
      disabled={disabled}
      onClick={onClick}
      style={{
        cursor: "pointer",
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        right: "1%",
        opacity: disabled ? "0" : "1",
        userSelect: "none",
        marginLeft: marginLeft,
        marginRight: marginRight
      }}
    >
      {children}
    </button>
  );
}

export function LeftArrow() {
    const { isFirstItemVisible, scrollPrev } =
      React.useContext(VisibilityContext);
  
    return (
      <Arrow disabled={isFirstItemVisible} onClick={() => scrollPrev()} marginRight='2px'>
        Left
      </Arrow>
    );
  }
  
  export function RightArrow() {
    const { isLastItemVisible, scrollNext } = React.useContext(VisibilityContext);
  
    return (
      <Arrow disabled={isLastItemVisible} onClick={() => scrollNext()} marginLeft='2px'>
        Right
      </Arrow>
    );
  }
