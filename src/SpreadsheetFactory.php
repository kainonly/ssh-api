<?php

namespace scf_extra;

use PhpOffice\PhpSpreadsheet\Spreadsheet;

class SpreadsheetFactory
{
    private $spreadsheet;

    public function __construct()
    {
        $this->spreadsheet = new Spreadsheet();
    }
}