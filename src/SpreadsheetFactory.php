<?php

namespace scf_extra;

class SpreadsheetFactory
{
    private $spreadsheet;

    public function __construct()
    {
        $this->spreadsheet = new Spreadsheet();
    }
}