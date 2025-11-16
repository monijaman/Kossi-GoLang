import json

# Load existing bankspec.json
with open('init-db/seeders/bankspec.json', 'r', encoding='utf-8') as f:
    data = json.load(f)

# Template
template = {
    "Bank Name": "",
    "Bank Name (Bangla)": "",
    "Brand": "",
    "Type": "Private Commercial Bank",
    "Established Year": "",
    "Headquarters": "",
    "SWIFT Code": "",
    "Routing Number": "",
    "License Type": "",
    "Ownership": "",
    "Chairman": "",
    "Managing Director / CEO": "",
    "Total Branches": "",
    "Total ATMs": "",
    "Total Agents": "",
    "Core Banking System": "",
    "Internet Banking": "Yes",
    "Mobile Banking App": "",
    "Mobile Banking App Name": "",
    "SMS Banking": "",
    "Phone Banking": "",
    "Debit Card": "Yes",
    "Credit Card": "Yes",
    "International Card Support": "",
    "Deposit Schemes": "",
    "Loan Schemes": "",
    "Islamic Banking Window": "",
    "Foreign Exchange Service": "",
    "Remittance Service": "",
    "Corporate Banking": "",
    "SME Banking": "",
    "Agricultural Loan Schemes": "",
    "Student Banking": "",
    "Women Banking": "",
    "Agent Banking Service": "",
    "Digital Wallet": "",
    "UPI / QR Payment Support": "",
    "ATM Network Partnership": "",
    "Visa / Mastercard / UnionPay Support": "",
    "Government Payment Support": "",
    "Utility Bill Payment Support": "",
    "Customer Care Phone": "",
    "Customer Care Email": "",
    "Website": "",
    "Facebook Page": "",
    "Head Office Address": "",
    "Helpline Availability": "",
    "Working Days": "",
    "Transaction Limit (Daily ATM)": "",
    "Transaction Limit (Daily App)": "",
    "Foreign Currency Account Support": "",
    "Nagad / bKash / Rocket / Upay Linked": ""
}

missing = {
    "AB Bank": {"Bank Name": "AB Bank Bangladesh Limited", "Bank Name (Bangla)": "এবি ব্যাংক বাংলাদেশ লিমিটেড", "Website": "https://www.abbank.com.bd", "Established Year": "1985", "SWIFT Code": "ABBABD6D"},
    "Bangladesh Commerce Bank": {"Bank Name": "Bangladesh Commerce Bank Limited", "Bank Name (Bangla)": "বাংলাদেশ কমার্স ব্যাংক লিমিটেড", "Website": "https://www.bcbl.com.bd", "Established Year": "1983", "SWIFT Code": "COMMBDDH"},
    "Bangladesh Development Bank": {"Bank Name": "Bangladesh Development Bank Limited", "Bank Name (Bangla)": "বাংলাদেশ উন্নয়ন ব্যাংক লিমিটেড", "Website": "https://www.bdbl.com.bd", "Established Year": "1973", "SWIFT Code": "BDBLBDDH", "Type": "Government Bank"},
    "Bangladesh Krishi Bank": {"Bank Name": "Bangladesh Krishi Bank", "Bank Name (Bangla)": "বাংলাদেশ কৃষি ব্যাংক", "Website": "https://www.krishibank.org.bd", "Established Year": "1972", "SWIFT Code": "KRISHBD6D", "Type": "Government Bank"},
    "Bank Alfalah": {"Bank Name": "Bank Alfalah Limited", "Bank Name (Bangla)": "ব্যাংক আল ফালাহ লিমিটেড", "Website": "https://www.balf.com", "Established Year": "1997", "SWIFT Code": "ALFAPH22"},
    "BASIC Bank": {"Bank Name": "Bangladesh Small Industries and Commerce Bank Limited", "Bank Name (Bangla)": "বিএএসআইসি ব্যাংক", "Website": "https://www.basicbank.org", "Established Year": "1989", "SWIFT Code": "BASICDH", "Type": "Government Bank"},
    "Citibank": {"Bank Name": "Citibank N.A", "Bank Name (Bangla)": "সিটিব্যাংক এন.এ.", "Website": "https://www.citibank.com.bd", "Established Year": "1972", "SWIFT Code": "CITIBDDX", "Type": "Foreign Bank"},
    "Citizens Bank": {"Bank Name": "Citizens Bank PLC", "Bank Name (Bangla)": "সিটিজেন্স ব্যাংক পিএলসি", "Website": "https://www.citizensbank-bd.com", "Established Year": "2005", "SWIFT Code": "CITIZBD6D"},
    "Commercial Bank of Ceylon": {"Bank Name": "Commercial Bank of Ceylon Limited", "Bank Name (Bangla)": "কমার্শিয়াল ব্যাংক অফ সিলন", "Website": "https://www.combank.lk", "Established Year": "1859", "SWIFT Code": "CEYLIN2C", "Type": "Foreign Bank"},
    "Community Bank": {"Bank Name": "Community Bank Bangladesh PLC", "Bank Name (Bangla)": "কমিউনিটি ব্যাংক বাংলাদেশ", "Website": "https://www.communitybankbd.com", "Established Year": "2000", "SWIFT Code": "COMMUNBD"},
    "Eastern Bank": {"Bank Name": "Eastern Bank PLC", "Bank Name (Bangla)": "ইস্টার্ন ব্যাংক পিএলসি", "Website": "https://www.ebl.com.bd", "Established Year": "1992", "SWIFT Code": "EBLKBDDH"},
    "Global Islami Bank": {"Bank Name": "Global Islami Bank PLC", "Bank Name (Bangla)": "গ্লোবাল ইসলামী ব্যাংক পিএলসি", "Website": "https://www.globalislamic.com.bd", "Established Year": "1999", "SWIFT Code": "GLBL2BD6", "Type": "Islamic Bank"},
    "Habib Bank": {"Bank Name": "Habib Bank Limited", "Bank Name (Bangla)": "হাবিব ব্যাংক লিমিটেড", "Website": "https://www.habib-bank.com", "Established Year": "1947", "SWIFT Code": "HABIBPK2X", "Type": "Foreign Bank"},
    "HSBC Bank": {"Bank Name": "The Hongkong and Shanghai Banking Corporation Limited", "Bank Name (Bangla)": "এইচএসবিসি ব্যাংক", "Website": "https://www.hsbc.com.bd", "Established Year": "1973", "SWIFT Code": "HSBKBDDH", "Type": "Foreign Bank"},
    "ICB Islamic Bank": {"Bank Name": "ICB Islamic Bank Limited", "Bank Name (Bangla)": "আইসিবি ইসলামী ব্যাংক লিমিটেড", "Website": "https://www.icb-islamicbank.com", "Established Year": "2004", "SWIFT Code": "ICIB2BD6", "Type": "Islamic Bank"},
    "Janata Bank": {"Bank Name": "Janata Bank PLC", "Bank Name (Bangla)": "জনতা ব্যাংক পিএলসি", "Website": "https://www.jbl.gov.bd", "Established Year": "1972", "SWIFT Code": "JANATBDD", "Type": "Government Bank"},
    "Mercantile Bank": {"Bank Name": "Mercantile Bank PLC", "Bank Name (Bangla)": "মার্চেন্টাইল ব্যাংক পিএলসি", "Website": "https://www.mercbank.com.bd", "Established Year": "1981", "SWIFT Code": "MERCBDDH"},
    "Modhumoti Bank": {"Bank Name": "Modhumoti Bank PLC", "Bank Name (Bangla)": "মধুমতি ব্যাংক পিএলসি", "Website": "https://www.modhumoticredit.com", "Established Year": "1999", "SWIFT Code": "MODHBDDH"},
    "Mutual Trust Bank": {"Bank Name": "Mutual Trust Bank PLC", "Bank Name (Bangla)": "মিউচুয়াল ট্রাস্ট ব্যাংক পিএলসি", "Website": "https://www.mutualtrustbank.com", "Established Year": "1997", "SWIFT Code": "MTBLBDDH"},
    "National Bank of Pakistan": {"Bank Name": "National Bank of Pakistan", "Bank Name (Bangla)": "ন্যাশনাল ব্যাংক অফ পাকিস্তান", "Website": "https://www.nbp.com.pk", "Established Year": "1949", "SWIFT Code": "NBPAPK3N", "Type": "Foreign Bank"},
    "NCC Bank": {"Bank Name": "NCC Bank PLC", "Bank Name (Bangla)": "এনসিসি ব্যাংক পিএলসি", "Website": "https://www.nccbank.com.bd", "Established Year": "1999", "SWIFT Code": "NCCBBDDH"},
    "NRBC Bank": {"Bank Name": "NRBC Bank PLC", "Bank Name (Bangla)": "এনআরবিসি ব্যাংক পিএলসি", "Website": "https://www.nrbc.com.bd", "Established Year": "2013", "SWIFT Code": "NRBC2BDD"},
    "Padma Bank": {"Bank Name": "Padma Bank PLC", "Bank Name (Bangla)": "পদ্মা ব্যাংক পিএলসি", "Website": "https://www.padmabank.com.bd", "Established Year": "2000", "SWIFT Code": "PADMBDDH"},
    "Probashi Kallyan Bank": {"Bank Name": "Probashi Kallyan Bank", "Bank Name (Bangla)": "প্রবাসী কল্যাণ ব্যাংক", "Website": "https://www.probashibank.gov.bd", "Established Year": "2001", "SWIFT Code": "PROBBD6D", "Type": "Government Bank"},
    "Rajshahi Krishi Unnayan Bank": {"Bank Name": "Rajshahi Krishi Unnayan Bank", "Bank Name (Bangla)": "রাজশাহী কৃষি উন্নয়ন ব্যাংক", "Website": "https://www.rkub.gov.bd", "Established Year": "1987", "SWIFT Code": "RKUBBD6D", "Type": "Government Bank"},
    "Shimanto Bank": {"Bank Name": "Shimanto Bank Limited", "Bank Name (Bangla)": "শিমান্ত ব্যাংক লিমিটেড", "Website": "https://www.shimantobkash.com", "Established Year": "2009", "SWIFT Code": "SHIMBD6D", "Customer Care Phone": "16477"},
    "Social Islami Bank": {"Bank Name": "Social Islami Bank PLC", "Bank Name (Bangla)": "সোশ্যাল ইসলামী ব্যাংক পিএলসি", "Website": "https://www.socialislamicbank.com", "Established Year": "2011", "SWIFT Code": "SOCIBDD6", "Type": "Islamic Bank"},
    "Standard Bank": {"Bank Name": "Standard Bank PLC", "Bank Name (Bangla)": "স্ট্যান্ডার্ড ব্যাংক পিএলসি", "Website": "https://www.standardbank.com.bd", "Established Year": "1985", "SWIFT Code": "STANBD6D"},
    "Standard Chartered Bank": {"Bank Name": "Standard Chartered Bank", "Bank Name (Bangla)": "স্ট্যান্ডার্ড চার্টার্ড ব্যাংক", "Website": "https://www.sc.com/bd", "Established Year": "1971", "SWIFT Code": "SCBLBDDH", "Type": "Foreign Bank"},
    "State Bank of India": {"Bank Name": "State Bank of India", "Bank Name (Bangla)": "স্টেট ব্যাংক অফ ইন্ডিয়া", "Website": "https://www.sbi.co.in", "Established Year": "1955", "SWIFT Code": "SBININBB", "Type": "Foreign Bank"},
    "United Commercial Bank": {"Bank Name": "United Commercial Bank PLC", "Bank Name (Bangla)": "ইউনাইটেড কমার্শিয়াল ব্যাংক পিএলসি", "Website": "https://www.ucbl.com.bd", "Established Year": "1983", "SWIFT Code": "UCBLBDDH"},
    "Uttara Bank": {"Bank Name": "Uttara Bank PLC", "Bank Name (Bangla)": "উত্তরা ব্যাংক পিএলসি", "Website": "https://www.uttarabank.gov.bd", "Established Year": "1972", "SWIFT Code": "UTTABDDH", "Type": "Government Bank"},
    "Woori Bank": {"Bank Name": "Woori Bank Bangladesh", "Bank Name (Bangla)": "উরি ব্যাংক বাংলাদেশ", "Website": "https://www.wooribank.com.bd", "Established Year": "2006", "SWIFT Code": "WOORIBDD", "Type": "Foreign Bank", "Customer Care Phone": "16477"}
}

# Add missing
for brand, details in missing.items():
    if not any(b.get('Brand') == brand for b in data):
        entry = template.copy()
        entry['Brand'] = brand
        for k, v in details.items():
            if v:
                entry[k] = v
        data.append(entry)

# Write back
with open('init-db/seeders/bankspec.json', 'w', encoding='utf-8') as f:
    json.dump(data, f, ensure_ascii=False, indent=2)

print(f"✅ Added missing banks. Total now: {len(data)}")
