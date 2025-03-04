export interface emailData{
    _source: {
        ID: number;
        "Message-ID": string;
        Date: string;
        from: string;
        to: string;
        subject: string;
        "Mime-Version": string;
        "Content-Type": string;
        "Content-Transfer-Encoding": string;
        "X-From": string;
        "X-To": string;
        "X-cc": string;
        "X-bcc": string;
        "X-Folder": string;
        "X-Origin": string;
        "X-FileName": string;
        Message: string;
    }
}