import React, { Component } from 'react'

export default class Element 
{
    id: string;
    htmlRender = () => {<></>};

    constructor(id: string)
    {
        this.id = id;
    }
}  