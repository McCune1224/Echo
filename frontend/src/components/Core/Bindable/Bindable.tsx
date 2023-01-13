import { useState } from 'react'

export default class Bindable<T>
{
    value: T; _SetValue: Function; 
    
    ChangeValue = (newValue: T) =>
    {
        this._SetValue(newValue);
        this.InvokeOnChange();
    }

    InvokeOnChange!:Function;
    
    constructor(defaultValue: T)
    {
        [this.value, this._SetValue] = useState(defaultValue);
        this.InvokeOnChange = () => {};
    }
}
