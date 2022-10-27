from docxtpl import DocxTemplate, RichText


def main():
    doc = DocxTemplate('demo.docx')
    ctx = {
        'title': '这是标题。',
        'name': '张三',
        'dept': RichText('真理部', color='#eecaf2'),
        'num': 18
    }
    doc.render(ctx)
    doc.save('out.docx')


if __name__ == '__main__':
    main()
