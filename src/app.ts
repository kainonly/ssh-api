import { utils, WorkBook, WorkSheet, write } from 'xlsx';

const App = () => {
  const workBook: WorkBook = utils.book_new();
  const workSheet: WorkSheet = utils.json_to_sheet([
    { A: 'S', B: 'h', C: 'e', D: 'e', E: 't', F: 'J', G: 'S' },
    { A: 1, B: 2, C: 3, D: 4, E: 5, F: 6, G: 7 },
    { A: 2, B: 3, C: 4, D: 5, E: 6, F: 7, G: 8 },
  ], { header: ['A', 'B', 'C', 'D', 'E', 'F', 'G'], skipHeader: true });
  utils.book_append_sheet(workBook, workSheet);

  write(workBook, {
    bookType: 'xlsx',
    bookSST: false,
    type: 'buffer',
  });

};

export { App };
