import React, { Component } from 'react'
import Element from './Element/Element';
import Bindable from './Bindable/Bindable';

export default class InputField extends Element
{
    activeState: Bindable<boolean>;
    valueHolder: Bindable<string>;
    
    constructor(id: string)
    {
        super(id);
        this.valueHolder = new Bindable<string>("");
        this.activeState = new Bindable<boolean>(true);
    }

    htmlRender = () =>
    { 
        const OnClick = () =>
        {
            this.activeState.ChangeValue(!this.activeState.value);
        }

        return(
            <div>
                <input type="text" disabled={!this.activeState.value} onChange={(event) => this.valueHolder.ChangeValue(event.target.value)}></input>
                <button onClick={OnClick}>
                    {
                        this.activeState.value?
                            <span>
                                <i> Active </i>
                            </span> :
                            <span>
                                <i>InActive</i>
                            </span>
                    }
                </button>
            </div>
        )
    } 
} 